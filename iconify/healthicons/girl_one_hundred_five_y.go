package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlOneHundredFiveY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 17a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm10.508 4.422a2 2 0 1 0-1.015-3.869c-3.809.999-6.672 1.446-9.485 1.434c-2.819-.011-5.69-.483-9.523-1.44a2 2 0 0 0-.97 3.881c1.85.462 3.535.828 5.14 1.09c-.316 2.1-1.996 5.559-2.965 7.425c-.324.626.066 1.395.764 1.484c5.937.752 9.692.772 15.155.012c.68-.095 1.06-.838.754-1.453c-.944-1.891-2.64-5.471-3-7.443c1.604-.26 3.291-.635 5.145-1.12ZM19 33.014V36.5a1.5 1.5 0 0 0 2.948.391l.94-3.483A38.923 38.923 0 0 1 19 33.014Zm7.11.387l.942 3.49A1.5 1.5 0 0 0 30 36.5V33a40.878 40.878 0 0 1-3.89.4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}