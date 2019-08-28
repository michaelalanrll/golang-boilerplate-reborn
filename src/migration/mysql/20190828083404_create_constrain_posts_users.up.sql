ALTER TABLE `posts` 
ADD CONSTRAINT `posts_users` 
FOREIGN KEY(user_id) 
REFERENCES users(id)