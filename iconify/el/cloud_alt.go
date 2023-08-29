package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.635 0 0 268.635 0 600s268.635 600 600 600s600-268.635 600-600S931.365 0 600 0zm5.347 311.938c112.641-2.458 224.208 89.791 256.494 210.944c128.177-12.623 199.88 157.997 101.366 244.557H224.7c-94.99-161.431 22.206-376.261 193.577-334.58C469.8 349 537.753 313.412 605.33 311.936l.017.002z"/>`),
		g.Group(children),
	)
}