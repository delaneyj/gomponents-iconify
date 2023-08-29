package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reflecth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.57 20.273h.854v-1.705h-.854v1.705zm0 3.413h.854V21.98h-.854v1.706zm0 3.41h.854V25.39h-.854v1.706zm0 2.593h.854v-.89h-.854v.89zm0-12.825h.854V15.16h-.854v1.705zm0-13.64h.854V1.52h-.854v1.705zm0 3.41h.854V4.93h-.854v1.705zm0 3.41h.854V8.34h-.854v1.705zm0 3.41h.854V11.75h-.854v1.705zm2.84-10.128V25.46h12.015L18.41 3.327zm.854 3.353l9.73 17.93h-9.73V6.68zm-5.73 18.78V3.327L1.522 25.46h12.015z"/>`),
		g.Group(children),
	)
}