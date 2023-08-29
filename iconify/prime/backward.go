package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.3 4.31a.756.756 0 0 0-.81.14l-7.27 6.82V5c0-.3-.18-.57-.45-.69a.756.756 0 0 0-.81.14l-7.47 7c-.15.14-.24.34-.24.55s.09.41.24.55l7.47 7a.763.763 0 0 0 .81.14c.27-.12.45-.39.45-.69v-6.27l7.27 6.82a.763.763 0 0 0 .81.14c.27-.12.45-.39.45-.69V5c0-.3-.18-.57-.45-.69Zm-9.58 12.96L5.1 12l5.62-5.27v10.54Zm8.53 0L13.63 12l5.62-5.27v10.54Z"/>`),
		g.Group(children),
	)
}