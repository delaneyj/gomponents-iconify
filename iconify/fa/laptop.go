package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M416 1024q-66 0-113-47t-47-113V160q0-66 47-113T416 0h1088q66 0 113 47t47 113v704q0 66-47 113t-113 47H416zm-32-864v704q0 13 9.5 22.5T416 896h1088q13 0 22.5-9.5t9.5-22.5V160q0-13-9.5-22.5T1504 128H416q-13 0-22.5 9.5T384 160zm1376 928h160v96q0 40-47 68t-113 28H160q-66 0-113-28t-47-68v-96h1760zm-720 96q16 0 16-16t-16-16H880q-16 0-16 16t16 16h160z"/>`),
		g.Group(children),
	)
}