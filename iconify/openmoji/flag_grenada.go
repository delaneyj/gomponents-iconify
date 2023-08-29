package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagGrenada(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="m12 48l24-12l-24-12v24zm48 0L36 36l24-12v24z"/><circle cx="36" cy="36" r="5" fill="#d22f27"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="m36 33.5l1.545 5l-4.045-3.09h5l-4.045 3.09l1.545-5z"/><path fill="#d22f27" d="M5 17v38h62V17Zm55 31H12V24h48Z"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="m35.105 22.242l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707zm10 0l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707zm-20 0l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707zm10 30l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707zm10 0l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707zm-20 0l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707z"/><g stroke-linecap="round" stroke-linejoin="round"><path fill="#fcea2b" stroke="#fcea2b" d="M18.63 37.917a6.781 6.781 0 0 0-1.257-2.622a6.781 6.781 0 0 0 2.621 1.258a1.767 1.767 0 0 0-1.561-2.125l-.075-.003l.002-.002h-.147a5.774 5.774 0 0 1-2.09-.377a5.776 5.776 0 0 1 .377 2.097v.139l.002-.002a1.763 1.763 0 0 0 2.128 1.637Z"/><path fill="#d22f27" stroke="#d22f27" d="M18.801 36.277a1.77 1.77 0 0 1 .447.893a1.578 1.578 0 0 1-1.34-1.34a1.77 1.77 0 0 1 .893.447Z"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}