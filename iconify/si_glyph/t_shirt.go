package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.021 1.062c-.917 0-1.705-.27-2.04-1.062c-.206.012-.735.033-.935.067c.383 1.088 1.752 1.873 2.975 1.873s2.58-.785 2.963-1.873c-.199-.034-.719-.056-.922-.067c-.334.793-1.123 1.062-2.041 1.062z"/><path d="M13.396.289c-.82 1.62-2.479 2.742-4.412 2.742c-1.924 0-3.572-1.11-4.398-2.717C2.33 1.33 1.043 3.83 1.043 6.466v.486h2.988v7.983h9.948V6.952H17v-.486c0-2.663-1.306-5.183-3.604-6.177z"/></g>`),
		g.Group(children),
	)
}