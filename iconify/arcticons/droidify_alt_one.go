package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroidifyAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.158 21.64h6.084a1 1 0 0 0 1-1v-1.083A12.508 12.508 0 0 0 24.229 6.91a12.242 12.242 0 0 0-12.47 12.24v1.489a1 1 0 0 0 1 1h6.083"/><rect width="24.484" height="16.547" x="11.758" y="26.953" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.04 26.953V24.6a2.96 2.96 0 1 1 5.92 0v2.354m-1.555 11.599a3.61 3.61 0 1 0-2.849-.017m7.86-29.812L34.64 4.5"/><circle cx="27.61" cy="15.115" r="1.211" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.584 8.724L13.36 4.5"/><circle cx="20.39" cy="15.115" r="1.211" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}