package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FuneralUrn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#f4aa41" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m44.797 15.175l3.362-2.94a2.12 2.12 0 0 0-.545-3.533a27.541 27.541 0 0 0-11.338-2.339a27.541 27.541 0 0 0-11.337 2.34a2.12 2.12 0 0 0-.545 3.531l3.66 3.201c-4.847 3.043-9.185 8.356-9.185 14.416c0 10.739 8.561 35.786 8.561 35.786h17.14s8.561-25.047 8.561-35.786a17.096 17.096 0 0 0-8.334-14.676Z"/><path fill="none" stroke="#e27022" stroke-linecap="square" stroke-linejoin="round" stroke-width="2" d="M28.053 15.435a30.948 30.948 0 0 1 16.744-.26"/><path fill="none" stroke="#e27022" stroke-linejoin="round" stroke-width="4" d="M25.352 59.201h21.296"/><path fill="none" stroke="#e27022" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.645 29.522h32.71m-32.71 4.566h32.176"/><path fill="#e27022" d="M44.797 15.175h-1.142c0 1.093 2.627 8.437 2.627 14.676c0 10.739-8.561 35.786-8.561 35.786h6.849s8.561-25.047 8.561-35.786a17.096 17.096 0 0 0-8.334-14.676Z"/><path fill="#e27022" d="M44.797 8.116c2.455 3.004.466 2.303-1.142 7.06c3.517 2.504 5.206-4.784 5.206-4.784s-1.875-2.846-4.064-2.276Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m44.797 15.175l3.362-2.94a2.12 2.12 0 0 0-.545-3.533a27.541 27.541 0 0 0-11.338-2.339a27.541 27.541 0 0 0-11.337 2.34a2.12 2.12 0 0 0-.545 3.531l3.66 3.201c-4.847 3.043-9.185 8.356-9.185 14.416c0 10.739 8.561 35.786 8.561 35.786h17.14s8.561-25.047 8.561-35.786a17.096 17.096 0 0 0-8.334-14.676Z"/>`),
		g.Group(children),
	)
}