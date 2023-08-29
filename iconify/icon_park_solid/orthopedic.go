package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orthopedic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOrthopedic0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="m23.028 36l10.975 6.999V24c7.674-5.21 10.493-10.022 8.457-14.434c-3.055-6.619-8-6.182-11.453-4.564C28.704 6.08 27.035 9.093 26 14.04c-3.038-6.466-7.305-9.7-12.8-9.7c-8.242 0-8.906 9.724-7.725 12.105c1.181 2.381 2.298 3.666 7.526 7.554c-.075 11.563.405 17.782 1.442 18.656c1.9 1.472 4.762-.746 8.585-6.656Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M27 25v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOrthopedic0)"/>`),
		g.Group(children),
	)
}