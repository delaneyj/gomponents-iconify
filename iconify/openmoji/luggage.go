package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="m40.701 37.652l2.598 1.5l-1 1.732l-2.598-1.5l1-1.732m-.365-2.366a1 1 0 0 0-.867.5l-2 3.464a1 1 0 0 0 .366 1.366l4.33 2.5a1 1 0 0 0 1.366-.366l2-3.464a1 1 0 0 0-.366-1.366l-4.33-2.5a.994.994 0 0 0-.5-.135Zm-9.839-6.019a1.002 1.002 0 1 1-.497.135a.993.993 0 0 1 .497-.135m0-2a3 3 0 1 0 2.601 1.5a2.987 2.987 0 0 0-2.6-1.5Zm28 15.001a1 1 0 1 1-.497.135a.992.992 0 0 1 .497-.135m0-2a3 3 0 1 0 2.601 1.5a2.987 2.987 0 0 0-2.6-1.5Zm-25.198 7.384l1 1.732l-2.598 1.5l-1-1.732l2.598-1.5m.365-2.366a.995.995 0 0 0-.499.134l-4.33 2.5a1 1 0 0 0-.366 1.366l2 3.464a1 1 0 0 0 1.366.366l4.33-2.5a1 1 0 0 0 .366-1.366l-2-3.464a1 1 0 0 0-.867-.5ZM15 38.5v3h-2v-3h2m1-2h-4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1Z"/><rect width="54" height="34" x="9" y="22" fill="#a57939" rx="3"/><rect width="4" height="38" x="19" y="20" fill="#6a462f" rx="1"/><rect width="4" height="38" x="49" y="20" fill="#6a462f" rx="1"/><rect width="7" height="6" x="38" y="36.268" fill="#e67a94" rx="1" transform="rotate(30 41.501 39.266)"/><circle cx="30.5" cy="30.268" r="3" fill="#fcea2b"/><circle cx="58.5" cy="43.269" r="3" fill="#61b2e4"/><rect width="7" height="6" x="29" y="46.268" fill="#f4aa41" rx="1" transform="rotate(-30 32.5 49.268)"/><rect width="7" height="6" x="10.5" y="37" fill="#b1cc33" rx="1" transform="rotate(90 14 40)"/><path fill="#a57939" d="M33.5 22v-2a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v2h3v-3.8a2.2 2.2 0 0 0-2.2-2.2h-6.6a2.2 2.2 0 0 0-2.2 2.2V22Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linejoin="round" d="M33.5 22v-2a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v2h3v-3.8a2.2 2.2 0 0 0-2.2-2.2h-6.6a2.2 2.2 0 0 0-2.2 2.2V22Z"/><path stroke-miterlimit="10" d="M53 56h7a3 3 0 0 0 3-3V25a3 3 0 0 0-3-3h-7m-34 0h-7a3 3 0 0 0-3 3v28a3 3 0 0 0 3 3h7"/><path stroke-linecap="round" stroke-linejoin="round" d="m40.067 40.75l-1.732-1l2-3.464l3.031 1.75"/><path stroke-linecap="round" stroke-miterlimit="10" d="M31.5 32a2 2 0 1 1-2-3.463m28 12.999a2 2 0 1 1 2 3.465"/><path stroke-linecap="round" stroke-linejoin="round" d="m31.067 47.786l-1.732 1l2 3.464l3.031-1.75M12 39.5v-2h4V41"/><rect width="4" height="38" x="19" y="20" stroke-miterlimit="10" rx="1"/><rect width="4" height="38" x="49" y="20" stroke-miterlimit="10" rx="1"/><path stroke-miterlimit="10" d="M23 22h26m0 34H23"/></g>`),
		g.Group(children),
	)
}