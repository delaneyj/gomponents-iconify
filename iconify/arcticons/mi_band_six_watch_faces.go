package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiBandSixWatchFaces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-miterlimit="10" d="M24 8.455a5.429 5.429 0 0 1 5.435 5.44v20.112c0 3.012-2.426 5.441-5.435 5.441s-5.435-2.429-5.435-5.44V13.895A5.429 5.429 0 0 1 24 8.455Z"/><path d="M31.86 33.521c0 3.206.098 6.024-1.358 8.356c-1.358 2.332-4.95 2.623-5.823 2.623h-1.358c-.874 0-4.465-.292-5.823-2.623s-1.359-5.15-1.359-8.356V14.48c0-3.206-.097-6.024 1.359-8.356C18.856 3.791 22.448 3.5 23.32 3.5h1.358c.874 0 4.465.292 5.823 2.623c1.359 2.332 1.359 5.15 1.359 8.356V33.52Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.896 21.896h4.308v4.308h-4.308zm4.208-4.808h-4.208L24 12.88l2.104 4.208zm0 15.928c-.1 2.805-4.208 2.805-4.308 0c.1-2.805 4.308-2.805 4.308 0Z"/>`),
		g.Group(children),
	)
}