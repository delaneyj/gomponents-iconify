package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Likee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 10.27l-.595-.595c-4.32-4.322-11.327-4.322-15.648 0h0c-4.321 4.32-4.321 11.327 0 15.648l.594.595"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.673 25.921l.596-.595c4.33-4.33 4.33-11.348 0-15.678h0c-4.329-4.329-11.348-4.329-15.677 0l-.596.596L8.32 25.921l15.677 15.677l15.677-15.677h0Zm-26.128 5.226L37.21 7.482m5.226 5.225L18.771 36.373m-5.203-15.671l-7.974-7.974m5.216-5.216l7.974 7.973"/>`),
		g.Group(children),
	)
}