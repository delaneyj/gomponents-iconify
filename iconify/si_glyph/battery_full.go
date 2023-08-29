package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.446 3H6.534C6.279 3 6 3.301 6 3.58v9.902c0 .273.278.497.535.497h3.912c.254 0 .553-.224.553-.497V3.58c0-.279-.297-.58-.554-.58z"/><path d="M11.306 1.021H9.999V.005H7v1.016H5.77c-.936 0-1.694.766-1.694 1.709v11.562c0 .942.759 1.709 1.694 1.709h5.536c.936 0 1.694-.767 1.694-1.709V2.73c0-.944-.759-1.709-1.694-1.709zM12 14c0 .525-.494 1-1 1H6c-.506 0-1.011-.475-1.011-1V3c0-.523.505-1 1.011-1h5c.506 0 1 .476 1 1v11z"/></g>`),
		g.Group(children),
	)
}