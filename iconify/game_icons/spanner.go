package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M331.188 16.72c-40.712-.002-81.41 15.408-112.438 46.436c-43.866 43.864-56.798 107-38.813 162.25L17.03 388.312v25.75l170.22-170.218c2.75 5.84 5.847 11.555 9.344 17.094L17.03 440.5v51.78H64l181.875-181.874a158.498 158.498 0 0 0 17.03 9.438L90.44 492.28h27.03l164.75-164.75c55.182 17.85 118.21 4.884 162-38.905c41.415-41.414 54.998-99.91 41.282-152.813L380.22 241.125l-90.033-23.938l-23.968-90.03L371.53 21.843a161.459 161.459 0 0 0-40.342-5.125z"/>`),
		g.Group(children),
	)
}