package goease

import (
	"regexp"
	"strings"
	"testing"
)

func TestCreateHash(t *testing.T) {
	hashRX, err := regexp.Compile(`^\$argon2id\$v=19\$m=65536,t=1,p=[0-9]{1,4}\$[A-Za-z0-9+/]{22}\$[A-Za-z0-9+/]{43}$`)
	if err != nil {
		t.Fatal(err)
	}

	hash1, err := ArgonCreateHash("pa$$word", ArgonDefaultParams)
	if err != nil {
		t.Fatal(err)
	}

	if !hashRX.MatchString(hash1) {
		t.Errorf("hash %q not in correct format", hash1)
	}

	hash2, err := ArgonCreateHash("pa$$word", ArgonDefaultParams)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Compare(hash1, hash2) == 0 {
		t.Error("hashes must be unique")
	}
}

func TestComparePasswordAndHash(t *testing.T) {
	hash, err := ArgonCreateHash("pa$$word", ArgonDefaultParams)
	if err != nil {
		t.Fatal(err)
	}

	match, err := ArgonComparePasswordAndHash("pa$$word", hash)
	if err != nil {
		t.Fatal(err)
	}

	if !match {
		t.Error("expected password and hash to match")
	}

	match, err = ArgonComparePasswordAndHash("otherPa$$word", hash)
	if err != nil {
		t.Fatal(err)
	}

	if match {
		t.Error("expected password and hash to not match")
	}
}

func TestDecodeHash(t *testing.T) {
	hash, err := ArgonCreateHash("pa$$word", ArgonDefaultParams)
	if err != nil {
		t.Fatal(err)
	}

	params, _, _, err := ArgonDecodeHash(hash)
	if err != nil {
		t.Fatal(err)
	}
	if *params != *ArgonDefaultParams {
		t.Fatalf("expected %#v got %#v", *ArgonDefaultParams, *params)
	}
}

func TestCheckHash(t *testing.T) {
	hash, err := ArgonCreateHash("pa$$word", ArgonDefaultParams)
	if err != nil {
		t.Fatal(err)
	}

	ok, params, err := ArgonCheckHash("pa$$word", hash)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected password to match")
	}
	if *params != *ArgonDefaultParams {
		t.Fatalf("expected %#v got %#v", *ArgonDefaultParams, *params)
	}
}

func TestStrictDecoding(t *testing.T) {
	// "bug" valid hash: $argon2id$v=19$m=65536,t=1,p=2$UDk0zEuIzbt0x3bwkf8Bgw$ihSfHWUJpTgDvNWiojrgcN4E0pJdUVmqCEdRZesx9tE
	ok, _, err := ArgonCheckHash("bug", "$argon2id$v=19$m=65536,t=1,p=2$UDk0zEuIzbt0x3bwkf8Bgw$ihSfHWUJpTgDvNWiojrgcN4E0pJdUVmqCEdRZesx9tE")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected password to match")
	}

	// changed one last character of the hash
	ok, _, err = ArgonCheckHash("bug", "$argon2id$v=19$m=65536,t=1,p=2$UDk0zEuIzbt0x3bwkf8Bgw$ihSfHWUJpTgDvNWiojrgcN4E0pJdUVmqCEdRZesx9tF")
	if err == nil {
		t.Fatal("Hash validation should fail")
	}

	if ok {
		t.Fatal("Hash validation should fail")
	}
}

func TestVariant(t *testing.T) {
	// Hash contains wrong variant
	_, _, err := ArgonCheckHash("pa$$word", "$argon2i$v=19$m=65536,t=1,p=2$mFe3kxhovyEByvwnUtr0ow$nU9AqnoPfzMOQhCHa9BDrQ+4bSfj69jgtvGu/2McCxU")
	if err != ArgonErrIncompatibleVariant {
		t.Fatalf("expected error %s", ArgonErrIncompatibleVariant)
	}
}