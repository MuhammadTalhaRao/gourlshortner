
# URL Shortner

This project creates a shortcode for URLs thats is unique for each URL based on hashing and checksum.
It has two features
1. Create ShortCode for URLs.
2. Redirect to original URL based on shortcode.

## Create shortcode
This api creates a unique shortcode for each url and reutrns it in response.

### Request Payload
{
    "url": "https://medium.com/@raotalha302.rt"
}

### Respones
{
    "shortCode": "AK_zpi6y_BNkjffTYaw6W3y9Vxcg_Mv8ZR_eXP_Ush8="
}

## Redirect URL
This api validates the shortcode then redirects user to the original URL.

### Request
It is a GET api so it takes data in param.

localhost:7088/AK_zpi6y_BNkjffTYaw6W3y9Vxcg_Mv8ZR_eXP_Ush8=

### Response 
It returns 301 http code and content of the original URL.

## How to run the code
Make sure you have opened the directory of the project in terminal.

### Liux
go run *.go

### Windows
go run .

## Test

There are 3 test sinarios 
1. Success case for Create shortcode.
2. Success case for fetching  and redirecting using shortcode.
3. Failure case for fetching and redirecting using shortcode.

To run the test, open the project directory in terminal and type following.

go test