package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M291 2H89L0 412h118l29-136h85q62 0 107.5-36t58.5-99q14-66-21-102.5T291 2zm-88 195h-39l25-112h58q31 0 40 29q-1 0-3.5-.5t-3.5-.5h-58zm84-56q-5 23-24 39t-41 17l14-61h52q0 3-1 5zm144 28q9-44-6-78q32 39 20 101q-14 63-59.5 98.5T278 326h-85l-29 136H46l5-22h99l29-136h85q62 0 107.5-36t59.5-99z"/>`),
		g.Group(children),
	)
}