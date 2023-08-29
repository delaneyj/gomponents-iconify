package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmbientMusicMod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.718" cy="28.956" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.998 43.96A11.069 11.069 0 0 0 3.14 29.225m14.008-.095a15.832 15.832 0 0 0-14.642-4.623M20.585 45.23a15.846 15.846 0 0 0-1.115-12.945m.668-6.88a20.48 20.48 0 0 0-14.467-5.957a20.739 20.739 0 0 0-2.726.179M25.484 45.45a20.566 20.566 0 0 0 .733-5.457a20.458 20.458 0 0 0-2.3-9.456m.379-14.473h5.606v3.68h-5.606z"/><circle cx="20.718" cy="28.933" r="3.578" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.296 19.744v9.189"/>`),
		g.Group(children),
	)
}