// Copyright (C) 2015 Constantin Schomburg <me@cschomburg.com>

package vdir

const TestVCard = `
BEGIN:VCARD
VERSION:4.0
N:Gump;Forrest;;;
FN:Forrest Gump
ORG:Bubba Gump Shrimp Co.
TITLE:Shrimp Man
PHOTO;MEDIATYPE=image/gif:http://www.example.com/dir_photos/my_photo.gif
TEL;TYPE=work,voice;VALUE=uri:tel:+1-111-555-1212
TEL;TYPE=home,voice;VALUE=uri:tel:+1-404-555-1212
ADR;TYPE=work;LABEL="100 Waters Edge\nBaytown, LA 30314\nUnited States of A
 merica":;;100 Waters Edge;Baytown;LA;30314;United States of America
ADR;TYPE=home;LABEL="42 Plantation St.\nBaytown, LA 30314\nUnited States of
  America":;;42 Plantation St.;Baytown;LA;30314;United States of America
EMAIL:forrestgump@example.com
REV:20080424T195243Z
END:VCARD
`
