package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneUpdateLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#solarSmartphoneUpdateLineDuotone0)"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13 2.001c3.114.01 4.765.108 5.828 1.17C20 4.344 20 6.23 20 10v4c0 3.771 0 5.657-1.172 6.828C17.657 22 15.771 22 12 22c-3.771 0-5.657 0-6.828-1.172C4 19.657 4 17.771 4 14v-2.999"/><path fill="currentColor" d="M2.73 4h-.75h.75Zm0 .8l-.554.506a.75.75 0 0 0 1.108 0L2.73 4.8Zm1.284-.294a.75.75 0 1 0-1.108-1.012l1.108 1.012Zm-1.46-1.012a.75.75 0 0 0-1.108 1.012l1.108-1.012ZM8.54 1.83A.75.75 0 0 0 9.71.892l-1.17.938ZM6.38-.75c-2.494 0-4.4 2.193-4.4 4.75h1.5c0-1.86 1.36-3.25 2.9-3.25v-1.5ZM1.981 4v.8h1.5V4h-1.5Zm1.303 1.306l.73-.8l-1.108-1.012l-.73.8l1.108 1.012Zm0-1.012l-.73-.8l-1.108 1.012l.73.8l1.108-1.012ZM9.71.892C8.914-.1 7.726-.75 6.381-.75v1.5c.843 0 1.617.404 2.159 1.08L9.71.892Z"/><path fill="currentColor" d="m10.27 3.2l.554-.505a.75.75 0 0 0-1.108 0l.554.505Zm-1.284.295a.75.75 0 1 0 1.108 1.01l-1.108-1.01Zm1.46 1.01a.75.75 0 0 0 1.108-1.01l-1.108 1.01Zm-5.87 1.801A.75.75 0 0 0 3.461 7.31l1.115-1.004ZM6.62 8.75c2.494 0 4.4-2.192 4.4-4.75h-1.5c0 1.861-1.36 3.25-2.9 3.25v1.5ZM11.02 4v-.8h-1.5V4h1.5ZM9.717 2.695l-.73.8l1.108 1.01l.73-.8l-1.108-1.01Zm0 1.01l.73.8l1.108-1.01l-.73-.8l-1.108 1.01ZM3.46 7.31c.79.878 1.908 1.44 3.158 1.44v-1.5c-.785 0-1.509-.35-2.043-.944L3.461 7.31Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M15 19H9" opacity=".5"/></g><defs><clipPath id="solarSmartphoneUpdateLineDuotone0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}