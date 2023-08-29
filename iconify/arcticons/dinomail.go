package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dinomail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.524 14.59S20.46 7.852 11.73 17.605m0-.001s-2.305.2-3.383-.738c-.643-.558-.223 3.5 1.975 5.418"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.342 19.98s-3.59 7.05.973 13.435m.001 0s.13 2.923-1.856 1.943c0 0 .293 2.809 1.165 3.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.624 38.848s-3.798-.788 1.674 4.063a7.654 7.654 0 0 1 5.535-3.103"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.558 41.824s1.473 3.514 8.09.367c0 0 3.844-.044 7.96-4.26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.648 42.192l5.99.177l1.234-2.064l1.15.08l3.585-2.25l-.717-1.611l2.694-2.686l-1.13-3.631l2.902-1.49l-1.493-4.344l2.942-4.22l-4.186-4.115l.207-5.826l-4.461.007l-3.154-5.718l-5.262 1.825l-2.998-1.455l-4.017 5.496l-3.493-.387l1.01 3.81l-2.014 1.541l.293 2.273"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.34 19.355a4.856 4.856 0 0 1-3.543-.906c-.634-.547-.117 2.665 2.047 4.462m-4.02 16.871s-.416-1.007.617-1.041"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M12.845 36.524a4.827 4.827 0 0 1-2.385-1.166m11.716 1.632c-1.641 0-3.304-.49-3.343-3.575c-.05-3.86 5.277-3.784 5.747-1.45c.816 3.967-2.528 4.454-1.793.652"/><ellipse cx="21.681" cy="33.865" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".917" ry="1.065"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M15.705 34.315c-.48-.584-1.201-.754-3.39-.9"/>`),
		g.Group(children),
	)
}