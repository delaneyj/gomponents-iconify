package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hammer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.216 14.949c.532.533.859.154 1.295-.281c.436-.436.815-.764.284-1.296c0 0-7.639-7.632-9.468-9.455L4.75 5.494l9.466 9.455zM2.048 7.015l.614-.604s-.271-.743.126-1.099s1.067-.136 1.067-.136L6.01 3.093s-.151-1.083.049-1.283c.2-.2 2.434-1.289 2.651-1.507l-.459-.459S5.123.219 4.784.558c-.199.2-1.689 1.704-2.751 2.766c0 0 .267.759-.083 1.109c-.351.351-1.141.099-1.141.099l-.623.623c-.263.265-.108.637.215.96l.686.686c.325.323.698.477.961.214z"/>`),
		g.Group(children),
	)
}