package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricCoffeePercolator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" fill-rule="evenodd" d="M49.5 21H42v2.5h7.5v4C49.5 30 44 41 44 41l3 6s5-15.5 5-19.5V21h-2.5Z" clip-rule="evenodd"/><path fill="#d0cfce" d="M18 27.5V21h4v4c0 2.5 4 6 4 6l-3 16s-5-15.5-5-19.5Z"/><path fill="#d0cfce" fill-rule="evenodd" d="M27.11 17c.635-2.838 3.926-5 7.89-5c3.964 0 7.255 2.162 7.89 5H43l5 33.5l-3 9.5H25l-3-9.5L27 17h.11Z" clip-rule="evenodd"/><path fill="#D0CFCE" fill-rule="evenodd" d="M32.667 60H30l-2-9.5L31.333 17H34l-3.333 33.5l2 9.5Z" clip-rule="evenodd"/><path fill="#3F3F3F" fill-rule="evenodd" d="M47.776 49H22.224L22 50.5l3 9.5h20l3-9.5l-.224-1.5Z" clip-rule="evenodd"/><circle cx="35.5" cy="55" r="2" fill="#D22F27" stroke="#9B9B9A" stroke-width="2"/><path fill="#3F3F3F" d="M33 8h4v4h-4z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M44 21h8v6.5c0 3.238-3.277 14.013-4.525 18M44 23.5h5.5v4c0 1.48-1.927 5.938-3.5 9.33m-23.5 8.59C21.234 41.371 18 30.717 18 27.5V21h4v4c0 1.667 1.778 3.778 2.963 5"/><path d="M43 17h-1c0 .05.004.099.011.148L43 17Zm-16 0l.989.148A.998.998 0 0 0 28 17h-1Zm-5 33.5l-.989-.148a.998.998 0 0 0 .035.45L22 50.5Zm3 9.5l-.954.301A1 1 0 0 0 25 61v-1Zm20 0v1a1 1 0 0 0 .954-.699L45 60Zm3-9.5l.954.301a.998.998 0 0 0 .035-.449L48 50.5ZM35 13c2.052 0 3.862.522 5.127 1.313C41.407 15.111 42 16.087 42 17h2c0-1.85-1.197-3.373-2.813-4.383C39.557 11.597 37.367 11 35 11v2Zm-7 4c0-.912.594-1.888 1.873-2.688C31.138 13.523 32.948 13 35 13v-2c-2.367 0-4.557.598-6.187 1.617C27.197 13.627 26 15.15 26 17h2Zm-1.989-.148l-5 33.5l1.978.296l5-33.5l-1.978-.296Zm-4.965 33.95l3 9.5l1.908-.603l-3-9.5l-1.908.602ZM25 61h10v-2H25v2Zm10 0h10v-2H35v2Zm10.954-.699l3-9.5l-1.908-.602l-3 9.5l1.908.602Zm3.035-9.949l-5-33.5l-1.978.296l5 33.5l1.978-.296Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M35.5 57a2 2 0 1 1 0-4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 50h21.5"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M27 17h13"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M33 8h4v4h-4z"/>`),
		g.Group(children),
	)
}