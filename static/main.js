$(document).ready(function () {
    function encode(r) { return r.replace(/[\x26\x0A\<>'"]/g, function (r) { return "&#" + r.charCodeAt(0) + ";" }) }

    $.getJSON("/api", function (data) {
        $(".players-public").text(data.public_player_count);
        $(".players-private").text(data.private_player_count);
        $(".players-total").text(data.total_player_count);

        $(".games-public").text(data.public_game_count);
        $(".games-private").text(data.private_game_count);
        $(".games-total").text(data.total_game_count);

        $(".in-progress-total").text(data.in_progress_count);
        $(".proxy-server-total").text(data.master_proxy_count);
    });

    $.getJSON("/api/public_games", function (data) {
        $(".public-games").empty();
        $.each(data, function () {
            $(".public-games").append(`
                <div class="bg-surface1 rounded-md shadow-md mb-4">
                    <div class="p-4 border-b border-surface0">
                        <div class="flex items-center">
                            <svg class="inline-block h-6 w-6 text-accent-error mr-2" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v8a2 2 0 01-2 2H4a2 2 0 01-2-2V6z" />
                            </svg>
                            <h3 class="text-lg font-medium">${this.game_name}</h3>
                            <span class="bg-surface0 text-white px-2 py-1 rounded-md text-sm ml-2">${this.title_id}</span>
                            <span class="bg-surface0 text-white px-2 py-1 rounded-md text-sm ml-2">v${this.title_version}</span>
                            
                            <svg class="inline-block h-5 w-5 mr-2 text-accent ml-auto" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z" />
                            </svg>
                            <span>${this.player_count}/${this.max_player_count}</span>
                        </div>            
                    </div>
                    <div class="p-4">
                        <blockquote class="mb-0">
                            <div class="flex items-center mb-2">
                                <svg class="inline-block h-5 w-5 mr-2 text-accent" viewBox="0 0 20 20" fill="currentColor">
                                    <path d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" />
                                </svg>
                                <span>${this.players.map(player => encode(player)).join(', <svg class="inline-block h-5 w-5 mr-2 text-accent" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 0010 16a5.986 5.986 0 004.546-2.084A5 5 0 0010 11z" clip-rule="evenodd" /></svg> ')}</span>
                            </div>
                            <footer class="font-normal text-accent flex items-center">
                                <svg class="inline-block h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                                    ${this.mode === "P2P" ? '<path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />' : '<path fill-rule="evenodd" d="M4 4a2 2 0 012-2h8a2 2 0 012 2v12a1 1 0 110 2h-3a1 1 0 01-1-1v-2a1 1 0 00-1-1H9a1 1 0 00-1 1v2a1 1 0 01-1 1H4a1 1 0 110-2V4zm3 1h2v4H7V5zm2 6a1 1 0 011-1h1a1 1 0 011 1v1a1 1 0 01-1 1h-1a1 1 0 01-1-1v-1zm5-1h1a1 1 0 011 1v1a1 1 0 01-1 1h-1a1 1 0 01-1-1v-1a1 1 0 011-1z" clip-rule="evenodd" />'}
                                </svg>
                                <span>${this.mode} (${this.status})</span>
                            </footer>
                        </blockquote>
                    </div>
                </div>
            `);
        });
    });
});