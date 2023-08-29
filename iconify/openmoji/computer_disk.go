package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="M11 10.958h50v50H11z"/><path fill="#D0CFCE" d="M52.55 47.117c-3.587 5.31-9.66 8.8-16.55 8.8c-11.023 0-19.958-8.936-19.958-19.959S24.977 16 36 16c6.857 0 12.906 3.458 16.5 8.725"/><circle cx="35.927" cy="35.887" r="19.958" fill="#D0CFCE"/><path fill="#D0CFCE" d="M17.513 26.386c3.078-5.938 8.954-10.29 16.097-11.121c11.43-1.331 21.773 6.855 23.103 18.284c1.33 11.43-6.856 21.773-18.284 23.103c-7.11.828-13.8-2.027-18.16-7.055"/><path fill="#D0CFCE" d="M52.55 47.046c-3.587 5.309-9.66 8.8-16.55 8.8c-11.023 0-19.958-8.936-19.958-19.96S24.977 15.93 36 15.93c6.857 0 12.906 3.457 16.5 8.724"/><circle cx="36" cy="35.958" r="5" fill="#3F3F3F"/><path fill="#D0CFCE" d="M44 36.501v5.474l17 2.458V28.5l-17 2.458v5.475"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M11 10.958h50v50H11z"/><path d="M52.55 47.046c-3.587 5.309-9.66 8.8-16.55 8.8c-11.023 0-19.958-8.936-19.958-19.96S24.977 15.93 36 15.93c6.857 0 12.906 3.457 16.5 8.724"/><circle cx="36" cy="35.958" r="5"/><path d="M44 36.501v5.474l17 2.458V28.5l-17 2.458v5.475z"/></g>`),
		g.Group(children),
	)
}