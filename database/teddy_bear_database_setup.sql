CREATE TABLE teddy_bears
(
    id SERIAL PRIMARY KEY,
    teddy_bear_name VARCHAR(100) NOT NULL,
    color VARCHAR(100) NOT NULL,
    occupation VARCHAR(100) NOT NULL,
    characteristic VARCHAR(100),
    age int
);


INSERT INTO
    teddy_bears
    (
    teddy_bear_name,
    color,
    occupation,
    characteristic,
    age
    )
VALUES
    (
        'Mr. Snuggles',
        'Brown',
        'Cop',
        'Welcoming',
        46
    ),
    (
        'Monkey',
        'Pink',
        'Chef',
        'Stern',
        8
    ),
    (
        'Teddy',
        'Gray',
        'Pharmisist',
        'Intelligent',
        21
    ),
    (
        'Dr. Evil',
        'Black',
        'Evil Scientist',
        'Evil',
        174
    ),
    (
        'Bob MacDonald',
        'White',
        'Farmer',
        'Caring',
        20
    ),
    (
        'Miss Penguin',
        'Purple',
        'Diner Owner',
        'Entrepreneurial',
        40
    ),
    (
        'Francis',
        'Black',
        'Lawyer',
        'Social',
        24
    )