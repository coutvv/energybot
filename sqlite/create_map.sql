CREATE TABLE IF NOT EXISTS CITY
(
    "id"     VARCHAR(2) NOT NULL PRIMARY KEY,
    "name"   TEXT,
    "region" VARCHAR(2)
);

INSERT OR IGNORE INTO CITY ("id", "name", "region")
VALUES ('1A', 'Flensburg', 'A'),
       ('2A', 'Kiel', 'A'),
       ('3A', 'Cuxhaven', 'A'),
       ('4A', 'Hamburg', 'A'),
       ('5A', 'Wilhelmshaven', 'A'),
       ('6A', 'Bremen', 'A'),
       ('7A', 'Hannover', 'A'),

       ('1B', 'Lubeck', 'B'),
       ('2B', 'Rostock', 'B'),
       ('3B', 'Schwerin', 'B'),
       ('4B', 'Torgelow', 'B'),
       ('5B', 'Berlin', 'B'),
       ('6B', 'Magdeburg', 'B'),
       ('7B', 'Frankfurt and der Oder', 'B'),

       ('1C', 'Osnabruck', 'C'),
       ('2C', 'Munster', 'C'),
       ('3C', 'Duisburg', 'C'),
       ('4C', 'Essen', 'C'),
       ('5C', 'Dortmund', 'C'),
       ('6C', 'Kassel', 'C'),
       ('7C', 'Dusseldorf', 'C'),

       ('1D', 'Halle', 'D'),
       ('2D', 'Lepzig', 'D'),
       ('3D', 'Erfurt', 'D'),
       ('4D', 'Dresden', 'D'),
       ('5D', 'Fulda', 'D'),
       ('6D', 'Wurzburg', 'D'),
       ('7D', 'Nurnberg', 'D'),

       ('1E', 'Aachen', 'E'),
       ('2E', 'Koln', 'E'),
       ('3E', 'Wiesbaden', 'E'),
       ('4E', 'Frankfurt-M', 'E'),
       ('5E', 'Trier', 'E'),
       ('6E', 'Saarbrucken', 'E'),
       ('7E', 'Mannheim', 'E'),

       ('1F', 'Stuttgart', 'F'),
       ('2F', 'Ausburg', 'F'),
       ('3F', 'Regensburg', 'F'),
       ('4F', 'Freiburg', 'F'),
       ('5F', 'Konstanz', 'F'),
       ('6F', 'Munchen', 'F'),
       ('7F', 'Passau', 'F');


CREATE TABLE IF NOT EXISTS CABLE
(
    "city_src"  VARCHAR(2) NOT NULL,
    "city_dest" VARCHAR(2) NOT NULL,
    "price"     INTEGER    NOT NULL,
    PRIMARY KEY (city_src, city_dest)
);

INSERT OR IGNORE INTO CABLE ("city_src", "city_dest", "price")
VALUES
    /*** REGION 'A' **/
    ('1A', '2A', 4),
    ('2A', '4A', 8),
    ('4A', '3A', 11),
    ('4A', '6A', 11),
    ('4A', '7A', 17),
    ('3A', '6A', 8),
    ('5A', '6A', 11),
    ('6A', '7A', 10),

    /*** CONN BW 'A' & 'B' */
    ('2A', '1B', 4),
    ('4A', '1B', 6),
    ('4A', '3B', 8),
    ('7A', '3B', 19),
    ('7A', '6B', 15),

    /*** CONN BW 'A' & 'C' */
    ('5A', '1C', 14),
    ('6A', '1C', 11),
    ('7A', '1C', 16),
    ('7A', '6C', 15),

    /*** CONN BW 'A' & 'D' */
    ('7A', '3D', 19),

    /*** REGION 'B' **/
    ('1B', '3B', 6),
    ('3B', '2B', 6),
    ('3B', '4B', 19),
    ('2B', '4B', 19),
    ('3B', '5B', 18),
    ('3B', '6B', 16),
    ('4B', '5B', 15),
    ('5B', '6B', 10),
    ('5B', '7B', 6),

    /*** CONN BW 'B' & 'D' */
    ('5B', '1D', 17),
    ('6B', '1D', 11),
    ('7B', '2D', 21),
    ('7B', '4D', 16),

    /*** REGION 'C' **/
    ('1C', '2C', 7),
    ('1C', '6C', 20),
    ('2C', '4C', 6),
    ('2C', '5C', 2),
    ('4C', '3C', 0),
    ('4C', '5C', 4),
    ('4C', '7C', 2),
    ('5C', '6C', 18),

    /*** CONN BW 'C' & 'D' */
    ('6C', '3D', 15),
    ('6C', '5D', 8),

    /*** CONN BW 'C' & 'E' */
    ('7C', '1E', 9),
    ('7C', '2E', 4),
    ('5C', '2E', 10),
    ('5C', '4E', 20),
    ('6C', '4E', 13),

    /*** REGION 'D' **/
    ('1D', '2D', 0),
    ('1D', '3D', 6),
    ('2D', '4D', 13),
    ('3D', '4D', 19),
    ('3D', '5D', 13),
    ('3D', '7D', 19),
    ('5D', '6D', 11),
    ('6D', '7D', 8),

    /*** CONN BW 'D' & 'E' */
    ('5D', '4E', 8),
    ('6D', '4E', 13),
    ('6D', '7E', 10),

    /*** CONN BW 'D' & 'F' */
    ('6D', '1F', 12),
    ('6D', '2F', 19),
    ('7D', '2F', 18),
    ('7D', '3F', 12),

    /*** REGION 'E' **/
    ('1E', '2E', 7),
    ('1E', '5E', 19),
    ('2E', '5E', 20),
    ('3E', '4E', 0),
    ('3E', '5E', 18),
    ('3E', '6E', 10),
    ('3E', '7E', 11),
    ('5E', '6E', 11),
    ('6E', '7E', 11),

    /*** CONN BW 'E' & 'F' */
    ('6E', '1F', 17),
    ('7E', '1F', 6),

    /*** REGION 'F' **/
    ('1F', '2F', 15),
    ('1F', '4F', 16),
    ('1F', '5F', 16),
    ('2F', '3F', 13),
    ('2F', '6F', 6),
    ('4F', '5F', 14),
    ('2F', '5F', 17),
    ('3F', '6F', 10),
    ('3F', '7F', 12),
    ('6F', '7F', 14);
