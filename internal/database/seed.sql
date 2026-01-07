INSERT INTO users (username) VALUES
('Harvey'), ('Jessica'), ('Donna');

INSERT INTO topics (name, created_by, description) VALUES 
('AMD', 1, 'This is the coolest topic ever.'),
('Arch', 2, 'I use Arch btw.'),
('ThinkPad', 3, 'Join the cult of enthusiasts because ThinkPad are truly the best laptops in the world!');

INSERT INTO posts (name, created_by, related_topic_id, content) VALUES 
('AMD Ryzen 7 is amazing!', 1, 1, 'The best CPU there is so far dont you think'),
('AMD has the best CPU', 3, 1, 'Just cuz I said so.'),
('AMD is superior', 2, 1, 'Superior to my company I mean'),
('Arch is amazing!', 1, 2, 'The best distro there is so far dont you think'),
('Arch has the best performance', 3, 2, 'Just cuz I said so.'),
('Arch is superior', 2, 2, 'I mean it.'),
('Thinkpads is amazing!', 1, 3, 'Thinkpad is the best laptop brand there is so far dont you think.'),
('Thinkpads have the best performance', 3, 3, 'Just cuz I said so.'),
('Thinkpads are superior', 2, 3, 'Thats a fact.');

INSERT INTO comments (created_by, related_post_id, content) VALUES 
(3, 1, 'I totally agree.'),
(2, 2, 'I totally agree.'),
(1, 3, 'I totally agree.'),
(3, 4, 'I totally agree.'),
(2, 5, 'I totally agree.'),
(1, 6, 'I totally agree.'),
(3, 7, 'I totally agree.'),
(2, 8, 'I totally agree.'),
(1, 9, 'I totally agree.');
