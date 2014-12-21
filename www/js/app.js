$(document).ready(function() {
    var app = Sammy("#main", function() {
        this.use('Mustache');
        this.get("#/", Controllers.index);
        this.get("#/library", Controllers.Library.index);
        this.get("#/library/playlists", Controllers.Library.playlists);
        this.get("#/music/now-playing", Controllers.Music.nowPlaying);
        this.get("#/ibus", Controllers.Ibus.index)

    });
    app.run("#/");
});

var Controllers = {
    index: function() {
        this.partial("/views/main.mustache");
    },
    Library: {
        index: function() {
            this.partial('/views/library/main.mustache')
        },
        playlists: function(context) {
            $.get('/library/playlists', function(data) {
                context.partial('/views/library/playlists-all.mustache', {playlists: data});
            });
        }
    },
    Music: {
        nowPlayinsdfg: function(context) {
            $.get("/jukebox/now-playing", function(data) {
                for (var c = 0; c < data.length; c++) {
                    var song = data[c];
                    var minutes = Math.floor(song.Time / 60);
                    var seconds = song.Time % 60;
                    if (seconds < 10 && seconds > 0) {
                        seconds = "0" + seconds;
                    } else if (seconds == 0) {
                        seconds = "00";
                    }
                    data[c].Length = minutes + ":" + seconds;
                }
                context.partial('/views/jukebox/now-playing.mustache', {songs: data});
            });
        },
        nowPlaying: function(context) {
            context.partial('/views/music/now-playing.mustache');
        }
    },
    Ibus: {
        index: function(context) {
            context.partial('/views/ibus/index.mustache')
        }
    }
}

var Playlists = {
    delete: function(playlist) {
        $.ajax('/jukebox/playlist/' + playlist, {
            type: 'DELETE'
        }).done(function() {
            $('#pl-' + playlist).remove();
        });
    },
    play: function(playlist) {
        $.ajax('/jukebox/playlist/' + playlist, {
            type: 'PUT'
        });
    }
}
var PlayerControls = {
    togglePlayPause: function() {
        $.post("/music/play", function() {

        });
    }
}