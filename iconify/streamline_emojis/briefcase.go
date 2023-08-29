package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M1.33 43.11a22.67 1.89 0 1 0 45.34 0a22.67 1.89 0 1 0-45.34 0Z" opacity=".15"/><path fill="#bf8256" d="M3.85 21.22h40.3V42H3.85Z"/><path fill="#915e3a" d="M42.89 21.22H5.11a1.27 1.27 0 0 0-1.26 1.26v6a1 1 0 0 0 .3.13l12.76 3.48a1.24 1.24 0 0 0 .33 0h13.52a1.24 1.24 0 0 0 .33 0l12.76-3.48a1 1 0 0 0 .3-.13v-6a1.27 1.27 0 0 0-1.26-1.26Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.85 21.22h40.3V42H3.85Z"/><path fill="#bf8256" d="M3.22 15.56V24a1.25 1.25 0 0 0 .93 1.21l12.76 3.48a1.26 1.26 0 0 0 .33.05h13.52a1.26 1.26 0 0 0 .33-.05l12.76-3.48a1.25 1.25 0 0 0 .93-1.21v-8.44a1.26 1.26 0 0 0-1.26-1.26h-39a1.26 1.26 0 0 0-1.3 1.26Z"/><path fill="#dea47a" d="M43.52 14.3h-39a1.26 1.26 0 0 0-1.26 1.26v4.25a1.26 1.26 0 0 1 1.26-1.26h39a1.26 1.26 0 0 1 1.26 1.26v-4.25a1.26 1.26 0 0 0-1.26-1.26Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.22 15.56V24a1.25 1.25 0 0 0 .93 1.21l12.76 3.48a1.26 1.26 0 0 0 .33.05h13.52a1.26 1.26 0 0 0 .33-.05l12.76-3.48a1.25 1.25 0 0 0 .93-1.21v-8.44a1.26 1.26 0 0 0-1.26-1.26h-39a1.26 1.26 0 0 0-1.3 1.26Z"/><path fill="#daedf7" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.07 26.89h13.85v3.78H17.07Zm3.15 3.78h7.56v2.52a1.26 1.26 0 0 1-1.26 1.26h-5a1.26 1.26 0 0 1-1.26-1.26v-2.52h0Z"/><path fill="#bf8256" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 24.37a3.78 3.78 0 0 0-3.78 3.78h7.56A3.78 3.78 0 0 0 24 24.37Z"/><path fill="#bf8256" d="M17.07 11.78v2.52h2.52v-2.52a1.26 1.26 0 0 1 1.26-1.26h6.3a1.26 1.26 0 0 1 1.26 1.26v2.52h2.52v-2.52A3.78 3.78 0 0 0 27.15 8h-6.3a3.78 3.78 0 0 0-3.78 3.78Z"/><path fill="#dea47a" d="M27.15 8h-6.3a3.78 3.78 0 0 0-3.78 3.78v1.73a3.78 3.78 0 0 1 3.78-3.78h6.3a3.78 3.78 0 0 1 3.78 3.78v-1.73A3.78 3.78 0 0 0 27.15 8Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.07 11.78v2.52h2.52v-2.52a1.26 1.26 0 0 1 1.26-1.26h6.3a1.26 1.26 0 0 1 1.26 1.26v2.52h2.52v-2.52A3.78 3.78 0 0 0 27.15 8h-6.3a3.78 3.78 0 0 0-3.78 3.78Z"/>`),
		g.Group(children),
	)
}