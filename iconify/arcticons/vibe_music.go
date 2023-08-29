package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibeMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.33 28.254v-5.248h1.896c2.206 0 2.223 5.248-.031 5.248h-1.866ZM9.67 23.006v5.248H7.775c-2.206 0-2.223-5.248.032-5.248H9.67Zm28.66 0c0-23.357-28.369-23.326-28.66 0m3.267-2.746v4.93m3.034-7.523v9.895m2.909-8.004v6.334m10.242 7.139v4.932m2.941-7.429v10.021m2.845-8.225v6.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.866 19.558l-.158 17.203c0 1.903-2.898 6.033-6.828 5.722c-5.606-.443-7.623-6.749-2.909-9.768c2.063-1.32 3.802-1.327 5.66-1.087l.315-20.442c1.996 2.58 4.73 4.322 7.682 6.165c5.192 3.24 7.247 8.152 3.136 11.254c-.238-2.2-.133-4.518-2.472-5.99c-1.437-1.04-2.956-2.158-4.426-3.057Z"/>`),
		g.Group(children),
	)
}