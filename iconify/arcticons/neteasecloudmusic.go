package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neteasecloudmusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.623 8.521c-3.557-7.516-12.102-3.778-9.383 5.285c1.347 4.49 6.057 8.04 5.986 12.727c-.062 4.104-4.063 6.993-8.035 6.74c-4.025-.254-6.542-1.96-7.28-6.093c-.596-3.337 2.592-6.872 5.554-8.52c6.268-3.488 16.126-.2 16.933 11.756c.652 9.674-11.21 16.153-22.487 11.594c-15.54-6.282-11.123-24.716-2.588-27.287"/>`),
		g.Group(children),
	)
}