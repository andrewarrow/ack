# ack
tcp simulation

# the playlist (2022)

https://www.themoviedb.org/tv/210195-the-playlist

In episode 4 of "The Playlist" (TV show drama about the history of Spotify) they
go into details about how they creatively broke the TCP rules and accepted
a little bit of packet loss to gain speed.

https://www.reddit.com/r/networking/comments/ydbyu3/a_tv_show_about_spotify_and_udp/

https://patents.google.com/patent/US20030079041A1/en

# simulation

This project's goal is to create a simulation of TCP in golang to demonstrate how
it all works.

# examples

```
./ack transfer --wire_speed=100
./ack transfer --wire_speed=100 --buffer_size=1000
./ack transfer --wire_speed=100 --buffer_size=1000 --process_speed=100
./ack transfer --wire_speed=10 --buffer_size=1000 --process_speed=100
```
