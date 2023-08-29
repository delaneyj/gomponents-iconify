package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaleTag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m44 278l160 192q21 24 49 24q25 0 43-15l139-117q20-18 23-45q3-25-15-49L283 76q-26-34-68-43L76 1q-20-6-38 9L21 25Q6 36 6 61l8 143q5 43 30 74zM68 42l141 34q22 5 44 27l160 192q8 8 5 18q0 6-9 17L270 447q-7 6-17 5.5t-15-7.5L78 253q-19-24-19-49L49 57zm57 70q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15zm11 149l196 62h7q14 0 21-15q3-8-1-16t-12-11l-196-62q-8-4-16.5 0T123 231q-6 23 13 30zm160 96q0 13-9 22.5t-23 9.5t-23-9.5t-9-22.5t9.5-22.5T264 325t22.5 9.5T296 357zm-64-170q0 13-9.5 22.5T200 219t-22.5-9.5T168 187t9-22.5t23-9.5t23 9.5t9 22.5z"/>`),
		g.Group(children),
	)
}