package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uterus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUterus0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" d="M32 18c-.923 6-3 11.5-7.5 11.5s-6-5-7.5-11.5s0-12 7.5-12s8.423 6 7.5 12Z"/><path stroke-linecap="round" d="M32.5 12.5c1.75-1.632 3.533-2.17 6.34-1.996c4.212.262 5.16 3.144 5.16 7.452m0 0c0 4.309-1.651 8.227-4.941 7.34c-3.29-.886-2.83-4.386-1.291-6.027c1.538-1.641 4.23-2.31 6.232-1.313ZM16.5 12c-1.77-1.544-3.86-1.67-6.413-1.495C6.257 10.767 4 13.692 4 18m0 0c0 4.308 2.897 8.182 5.888 7.296c2.992-.887 2.574-4.387 1.175-6.028C9.663 17.628 5.821 17.002 4 18Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 27v16m9-16v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUterus0)"/>`),
		g.Group(children),
	)
}