package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FF473E" d="M45.4 35.2V2.1c0-4.9-18.8 9.8-18.8 15.1v18h18.8z"/><path fill="#D1CFC3" d="M25.8 41.1c-.4 0-.7-.3-.7-.7v-7.2c0-.4.3-.7.7-.7h20.9c.4 0 .7.3.7.7v7.2c0 .4-.3.7-.7.7H25.8z"/><path fill="#BFBCAF" d="M50.1 39.6c0-.4-.3-.7-.7-.7H22.6c-.4 0-.7.3-.7.7v27.9C21.9 70 28.2 72 36 72s14.1-2 14.1-4.5V39.6z"/><ellipse cx="36" cy="9.4" fill="#FF6E83" rx="4" ry="12.6" transform="rotate(45.001 35.977 9.37)"/><path fill="#FC7570" d="M41.9 30.5c-.8 0-1.5-.7-1.5-1.5V17.5c0-.8.7-1.5 1.5-1.5s1.5.7 1.5 1.5V29c-.1.9-.7 1.5-1.5 1.5z"/><circle cx="36" cy="53.9" r="3.8" fill="#C4F0F2"/>`),
		g.Group(children),
	)
}