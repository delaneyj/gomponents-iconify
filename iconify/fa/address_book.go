package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1201 1238q0-57-5.5-107t-21-100.5t-39.5-86t-64-58t-91-22.5q-6 4-33.5 20.5T904 909t-40.5 20t-49 17t-46.5 5t-46.5-5t-49-17t-40.5-20t-42.5-24.5T556 864q-51 0-91 22.5t-64 58t-39.5 86t-21 100.5t-5.5 107q0 73 42 121.5t103 48.5h576q61 0 103-48.5t42-121.5zm-173-594q0-108-76.5-184T768 384t-183.5 76T508 644q0 107 76.5 183T768 903t183.5-76t76.5-183zm636 540v192q0 14-9 23t-23 9h-96v224q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V160Q0 94 47 47T160 0h1216q66 0 113 47t47 113v224h96q14 0 23 9t9 23v192q0 14-9 23t-23 9h-96v128h96q14 0 23 9t9 23v192q0 14-9 23t-23 9h-96v128h96q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}