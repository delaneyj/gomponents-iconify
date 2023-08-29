package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chesswalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.46 42.105c1.983 0 3.04-1.914 3.04-3.37s-.2-12.034-4.966-21.697c-4.767-9.662-13.631-6.41-13.631-6.41s-.529-4.697-1.914-4.626c-1.385.07-2.313 2.712-2.313 2.712l-1.855-.07s-2.184-2.842-3.44-2.783s-1.327 1.914-.599 5.495c0 0-.657 1.585-2.512 2.982s-.799 3.769-.998 4.825c-.2 1.057-2.783 4.967-3.64 6.082c-.857 1.115-1.984 3.698-.129 5.495a28.042 28.042 0 0 0 3.698 2.982s.928 1.327 2.912-1.127c1.984-2.454 2.783-4.896 7.08-7.338c3.111-1.914 3.839-1.714 3.839-1.714s.657 1.526.2 3.839c-.458 2.313-5.425 8.066-6.68 9.463s-2.713 4.297.128 5.295h21.78v-.035Z"/>`),
		g.Group(children),
	)
}