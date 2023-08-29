package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M20 20h4v4h-4zm0 28h4v4h-4zm28-28h4v4h-4z"/><g stroke="#000" stroke-linejoin="round" stroke-width="2"><path fill="#fff" d="M12 12h48v48H12z"/><path fill="#3F3F3F" d="M20 20h4v4h-4zm0 28h4v4h-4zm28-28h4v4h-4z"/></g><circle cx="18" cy="40" r="1"/><circle cx="16" cy="38" r="1"/><circle cx="20" cy="38" r="1"/><circle cx="34" cy="46" r="1"/><circle cx="40" cy="38" r="1"/><circle cx="40" cy="28" r="1"/><circle cx="32" cy="16" r="1"/><circle cx="46" cy="32" r="1"/><circle cx="52" cy="32" r="1"/><circle cx="52" cy="44" r="1"/><circle cx="54" cy="48" r="1"/><circle cx="56" cy="56" r="1"/><circle cx="32" cy="56" r="1"/><circle cx="44" cy="56" r="1"/><circle cx="46" cy="54" r="1"/><circle cx="44" cy="52" r="1"/><circle cx="16" cy="32" r="1"/><circle cx="40" cy="54" r="1"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M12 12h48v48H12z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M16 16h12v12H16z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M20 20h4v4h-4zm-4 24h12v12H16z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M20 48h4v4h-4zm24-32h12v12H44z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M48 20h4v4h-4z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 36v-2h8m-6 0v-2m4 2v6m0-2h2m12-6v-2m18 4h-2m-12 8h2v-2m-16-8h2m4 0h6m-2-16v4h-2v8m0-2h-4v2m4-8h-4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36 22h-2v-4m-6 18h8m-6 0v4h-2m6-4v2m-2 6v-2h6v6h4v-2h8v10"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36 40v4h6m4-4h-4v8h2m4-14v4h2v4h-2v4m2-8v-2h2m0 14h-4v2m-16 0h2v2h2v-4m20-18v6h-2m-10-2v-2m12 8v2m-2 8h2M40 22v2"/>`),
		g.Group(children),
	)
}