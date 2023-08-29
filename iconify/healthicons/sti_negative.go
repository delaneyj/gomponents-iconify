package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StiNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsStiNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm28 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-11.4 6a2 2 0 0 0-1.774 1.077l-2.6 5a2 2 0 0 0 .525 2.485l4.118 3.294c1.07.648 2.18.406 2.908-.174c.452-.36.7-.806.745-1.198c.04-.347-.062-.768-.568-1.2a130.977 130.977 0 0 0-3.394-2.793l-.05-.039l-.012-.01l-.003-.002l.151-.194L17.814 18H18v1.08c.174.139.426.34.74.595c.63.51 1.509 1.233 2.511 2.087c.549.468.919 1.01 1.114 1.586a4 4 0 1 1 2.77 7.488l.87 11.317A2 2 0 0 0 30 42V26.947c1.715-.19 3.153-.878 4.207-1.953c1.229-1.253 1.808-2.903 1.793-4.513A6.46 6.46 0 0 0 34.154 16c-1.244-1.259-3.014-2-5.135-2H16.6ZM18 28.238a4.54 4.54 0 0 0 2.085-.415a4.007 4.007 0 0 0 2.78 3.014l-.87 11.316A2 2 0 0 1 18 42V28.239ZM26 27a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm4-4.1v-4.792c.601.143 1.02.413 1.307.703A2.46 2.46 0 0 1 32 20.519c.006.64-.223 1.24-.65 1.674c-.28.287-.71.564-1.35.707Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsStiNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}