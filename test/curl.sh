echo 'GET /users/{email}'
curl -X 'GET' 'http://127.0.0.1:8080/users/user1@mail.com' \
  -H 'accept: application/json' \
  -i

echo
echo 'POST /users'
curl -X 'POST' 'http://127.0.0.1:8080/users' \
  -H 'accept: application/json' \
  -H 'content: application/json' \
  -d '{
    "email": "new@new.new",
    "password": "pass",
    "nickname": "nicko"
  }' \
  -i

echo
echo 'PATCH /users/{email}'
curl -X 'PATCH' 'http://127.0.0.1:8080/users/user1@mail.com' \
  -H 'accept: application/json' \
  -H 'content: application/json' \
  -d '{
    "password": "newPassword"
  }' \
  -i
