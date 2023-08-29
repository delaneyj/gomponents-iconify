package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBow0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6.544 14.262C11.403 15.23 16.603 18.754 20 23c0 10.63-7.356 13.237-12.771 13.834c-1.478.163-2.451-1.401-1.853-2.763C6.826 30.773 8 27.3 8 25c0-2.391-1.906-5.418-3.186-8.573c-.49-1.206.453-2.418 1.73-2.165Zm34.912 0C36.597 15.23 31.397 18.754 28 23c0 10.63 7.356 13.237 12.771 13.834c1.478.163 2.451-1.401 1.853-2.763C41.174 30.773 40 27.3 40 25c0-2.391 1.906-5.418 3.186-8.573c.49-1.206-.453-2.418-1.73-2.165Z"/><path fill="#555" d="M20 21h8v8h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBow0)"/>`),
		g.Group(children),
	)
}