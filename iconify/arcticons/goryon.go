package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goryon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m20.925 3l-2.545 8.488l-.207.07l-5.898-7.897s1.177 6.291 1.775 9.672c-5.386 3.145-8.651 8.616-8.655 14.497C5.395 37.313 13.724 45 24 45h0c10.275 0 18.605-7.686 18.605-17.169v0c-.003-7.06-4.687-13.398-11.808-15.976L32.64 7.95l-4.144 3.22a20.02 20.02 0 0 0-2.475-.403l-1.148-5.499l-1.984 5.468c-.1.005-.199.011-.298.018L20.925 3Z"/><path d="m28.701 33.73l-4.512 8.066l-4.737-7.937l9.25-.13Z"/><ellipse cx="16.822" cy="25.053" rx="4.728" ry="4.723"/><circle cx="17.706" cy="26.143" r="1.191"/><ellipse cx="31.094" cy="25.053" rx="4.728" ry="4.723"/><circle cx="30.209" cy="26.143" r="1.191"/></g>`),
		g.Group(children),
	)
}