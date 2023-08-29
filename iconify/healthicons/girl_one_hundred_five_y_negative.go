package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlOneHundredFiveYNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsGirl0105yNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm24 17a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm10.508 4.422a2 2 0 1 0-1.015-3.869c-3.809.999-6.672 1.446-9.485 1.434c-2.819-.011-5.69-.483-9.523-1.44a2 2 0 0 0-.97 3.881c1.85.462 3.535.828 5.14 1.09c-.316 2.1-1.996 5.559-2.965 7.425c-.324.626.066 1.395.764 1.484c5.937.752 9.692.772 15.155.012c.68-.095 1.06-.838.754-1.453c-.944-1.891-2.64-5.471-3-7.443c1.604-.26 3.291-.635 5.145-1.12ZM19 36.5v-3.486a38.92 38.92 0 0 0 3.889.394l-.94 3.483a1.5 1.5 0 0 1-2.949-.39Zm8.052.391l-.942-3.49A40.942 40.942 0 0 0 30 33v3.5a1.5 1.5 0 0 1-2.948.391Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsGirl0105yNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}