package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDropCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopWaterDropCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.025 7.402c-.37-.19-.468-.44-.44-.713c.027-.252.162-.525.294-.791l.072-.146c.218-.45.365-.866-.123-1.11a14.118 14.118 0 0 0-1.705.295c-4.115.981-6.731 3.637-7.71 5.531a7.133 7.133 0 1 0 9.612-3.066Zm-2.416-.24c-2.945.988-4.763 2.955-5.418 4.224a5.133 5.133 0 1 0 6.916-2.207a2.867 2.867 0 0 1-1.121-1a2.574 2.574 0 0 1-.378-1.017Z"/><path d="M10.458 13.183a1 1 0 0 0-1.965-.374l.982.191l-.982-.191v.002l-.001.003l-.001.006l-.003.017a2.453 2.453 0 0 0-.029.205a4.273 4.273 0 0 0-.023.506c.005.4.067.973.336 1.541c.269.566.667.975.97 1.231a4.198 4.198 0 0 0 .578.408l.015.009l.006.003l.003.001c.001.001.002.002.476-.88l-.474.881a1 1 0 0 0 .955-1.757l-.007-.004a2.196 2.196 0 0 1-.262-.19a1.737 1.737 0 0 1-.452-.558a1.741 1.741 0 0 1-.144-.71a2.273 2.273 0 0 1 .02-.331l.002-.009Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopWaterDropCircleFilled0)"/></g>`),
		g.Group(children),
	)
}