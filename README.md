# Forum

## Description

This project consists in creating a web forum that allows:
   * communication between users;
   * associating categories to posts;
   * liking and disliking posts and comments;
   * filtering posts.

### SQLite 
In order to store the data in the forum (like users, posts, comments, etc.) database library SQLite is used. 
SQLite is a popular choice as an embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

### Authentication 
In this segment the client is be able to register as a new user on the forum, by inputting their credentials. Login session was also created to access the forum and be able to add posts and comments. 
Cookies are used to allow each user to have only one opened session. Each of this sessions contains an expiration date. 

### Instructions for user registration:
   + Asks for email
   + When the email is already taken error response is returned
   + Asks for username
   + Asks for password

The forum is able to check if the email provided is present in the database and if all credentials are correct. It will check if the password is the same with the one provided and, if the password is not the same, it will return an error response.

### Communication 
In order for users to communicate between each other, they are able to create posts and comments.
   * Only registered users will be able to create posts and comments;
   * When registered users are creating a post they can associate one or more categories to it;
   * The posts and comments are visible to all users (registered or not);
   * Non-registered users sre only able to see posts and comments;

### Likes and Dislikes 
Only registered users are able to like or dislike posts and comments. 
The number of likes and dislikes are visible by all users (registered or not).

### Filter 
Filter mechanism allows users to filter the displayed posts by : 
   + categories 
   + created 
   + posts 
   + liked 
   + posts

The last two are only available for registered users and refers to the logged in user.

### Docker
Docker is used to be able to containerize the project.

To build docker image run the following command:
make build

To run container for docker image run the following command:
make d-run

### Run Locally
1. Run the following command: "make run" and click on the generated URL address to go to the web page
2. Sign up to create your own account
3. After authorization, sign in to go to your profile
4. Use any of desired available functions: post/comment/like/dislike 