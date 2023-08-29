package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indentleftalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M904 640H568q-50 0-85-37.5T448 512t35-90.5t85-37.5h336q50 0 85 37.5t35 90.5t-35 90.5t-85 37.5zm0-384H248q-50 0-85-37.5T128 128t35-90.5T248 0h656q50 0 85 37.5t35 90.5t-35 90.5t-85 37.5zm-712 70v122h160q13 0 22.5 9.5T384 480v64q0 13-9.5 22.5T352 576H192v121q-7 7-15.5 7t-15.5-7L6 529q-6-7-6-17.5T6 494l155-168q7-7 15.5-7t15.5 7zm56 442h656q50 0 85 37.5t35 90.5t-35 90.5t-85 37.5H248q-50 0-85-37.5T128 896t35-90.5t85-37.5z"/>`),
		g.Group(children),
	)
}