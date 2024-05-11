package goease

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

/*
	 GenerateNewJwtTokenHelper creates a new JWT token based on the provided claims and secret key.

	 This function is responsible for generating a JWT (JSON Web Token) using the specified claims
	 and a secret key. It uses the HMAC SHA256 signing method for token generation. The function is
	 mainly used for creating refresh tokens, but it is generic enough to be used for any JWT creation
	 where HMAC SHA256 is the appropriate signing method.

	 Example Usage:

		claims := jwt.MapClaims{
		    "sub": "1234567890",
		    "name": "John Doe",
		    "admin": true,
		    "iat": time.Now().Unix(),
		}
		secretKey := []byte("your-256-bit-secret")
		token, err := GenerateNewJwtTokenHelper(claims, secretKey)
		if err != nil {
		    fmt.Println("Error generating JWT token:", err)
		} else {
		    fmt.Println("Generated JWT token:", token)
		}

	 Parameters:

		claims: A jwt.Claims object containing the claims for the token. These claims are the
		        payload of the token and typically include user details and token metadata.
		secretKey: A byte slice representing the secret key used for signing the token.

	 Returns:

		A string representing the generated JWT token if the process is successful.
		An error if there is any issue in token generation, such as an error in signing.

	 Note:

		The function currently only supports HMAC SHA256 signing method. If other signing methods
		are required, additional functions or modifications to this function would be necessary.

	 Latest Modified: [Sat, 06 Jan 2024 03:51:24 GMT]
*/
func GenerateNewJwtTokenHelper(claims jwt.Claims, secretKey []byte) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return refreshTokenString, nil
}

type TokenClaims struct {
	Iss string // Issuer
	Sub string // Subject
	Aud string // Audience
	// minutes
	AccessExp  int64
	RefreshExp int64
}

/*
	GenerateDynamicJWTWithClaimsHelper creates an access token and a refresh token based on the provided claims.

This function takes two arguments: `tokenClaims` which is of type TokenClaims and contains the standard JWT claims like issuer (iss), subject (sub), audience (aud), and the expiration times for both access and refresh tokens. The second argument `additionalClaims` is a map of interface{} which allows adding extra information to the token.

The function performs the following operations:
1. It initializes the claims for the access token using both standard claims from `tokenClaims` and additional claims from `additionalClaims`.
2. It sets the "token_type" for the access token to "access".
3. It calls `GenerateNewJwtTokenHelper` to create the JWT access token.
4. It repeats similar steps for the refresh token, setting its "token_type" to "refresh".

Parameters:
- tokenClaims: TokenClaims - Struct containing standard JWT claims like issuer, subject, and audience, as well as expiration times for both tokens.
- additionalClaims: map[string]interface{} - Map containing additional claims to be included in the token.

Returns:
- string: The generated JWT access token.
- string: The generated JWT refresh token.
- error: An error message in case of failure in token generation.

Errors:
- If `GenerateNewJwtTokenHelper` fails to generate either the access or refresh token, the function returns an error.

Note:
- It's crucial to ensure that `GenerateNewJwtTokenHelper` and configs.JWT_SECRET are properly set up as they play a key role in token generation.
- The function assumes the `tokenClaims` struct is properly populated, especially the expiration times for both tokens
Latest Modified: [Sat, 06 Jan 2024 03:51:24 GMT]
*/
func GenerateDynamicJWTWithClaimsHelper(tokenClaims TokenClaims, additionalClaims map[string]interface{}, jwtSecret string) (string, string, error) {
	secret := []byte(jwtSecret)
	// Prepare accessTokenClaims by merging StandardClaims and additionalClaims
	accessTokenClaims := jwt.MapClaims{
		"iss": tokenClaims.Iss,
		"sub": tokenClaims.Sub,
		"aud": tokenClaims.Aud,
		"iat": time.Now().Unix(),
		"exp": tokenClaims.AccessExp,
	}

	// Adding additional claims for access token
	for key, value := range additionalClaims {
		accessTokenClaims[key] = value
	}
	accessTokenClaims["token_type"] = "access"

	accessTokenString, err := GenerateNewJwtTokenHelper(accessTokenClaims, secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshTokenClaims := jwt.MapClaims{
		"iss": tokenClaims.Iss,
		"sub": tokenClaims.Sub,
		"aud": tokenClaims.Aud,
		"iat": time.Now().Unix(),
		"exp": tokenClaims.RefreshExp,
	}

	for key, value := range additionalClaims {
		refreshTokenClaims[key] = value
	}
	refreshTokenClaims["token_type"] = "refresh"

	refreshTokenString, err := GenerateNewJwtTokenHelper(refreshTokenClaims, secret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

/*
	DecodeTokenHelper decodes and validates a JWT token string and returns its claims.

This function takes a JWT token as a string and decodes it to extract the claims. It also performs validation of the token to ensure its integrity and authenticity. The validation includes checking the signing method to ensure it matches the expected algorithm.

Parameters:
- tokenString: string - The JWT token that needs to be decoded and validated.

Returns:
- jwt.MapClaims: A map of claims (key-value pairs) extracted from the token if it is valid.
- error: An error message if the token is invalid or if any other error occurs during the decoding process.

Process:
1. The function uses `jwt.Parse` to parse the token string.
2. Inside the parsing function, it checks if the token's signing method matches the expected HMAC signing method.
  - If the signing method is not as expected, it returns an error.

3. If the signing method is correct, it returns the secret key used for signing the token.
4. After parsing, the function checks if the token is valid and if the claims type assertion is successful.
  - If successful, it returns the claims.
  - If not, it returns an error which could be due to an invalid token or a failure in the type assertion of claims.

Error Handling:
- The function returns an error if the token signing method is not HMAC.
- It also returns an error if the token is not valid or if the claims cannot be asserted as jwt.MapClaims.

Note:
- The secret key used for validating the token signature is retrieved from `configs.JWT_SECRET`.
- It's important that `configs.JWT_SECRET` is consistent with the secret key used for generating the tokens.
*/
func DecodeTokenHelper(tokenString string, jwtSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {

		return nil, err
	}
}
