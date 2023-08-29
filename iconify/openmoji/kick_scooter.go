package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KickScooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="18.825" cy="55.057" r="5" fill="#D0CFCE"/><circle cx="55.342" cy="55.057" r="5" fill="#D0CFCE"/><path fill="#EA5A47" d="M27.738 52.057h22.604v.583c0 1.657-1.223 3-2.73 3H30.467c-1.508 0-2.73-1.343-2.73-3v-.583z"/><path fill="#D0CFCE" d="m28.198 17.375l-1.058-.231a1 1 0 0 1-.763-1.19l.445-2.036a1 1 0 0 1 1.19-.763l1.059.232a1 1 0 0 1 .763 1.19l-.445 2.035a1 1 0 0 1-1.19.764z"/><path fill="#EA5A47" d="m41.048 15.864l-.25 1.324a.954.954 0 0 1-1.112.758l-4.132-.783a.954.954 0 0 1-.758-1.112l.25-1.323a.954.954 0 0 1 1.112-.758l4.132.782a.954.954 0 0 1 .758 1.112zm-19.5-3.951l-.251 1.325a.953.953 0 0 1-1.11.757l-4.12-.78a.953.953 0 0 1-.758-1.11l.251-1.326a.953.953 0 0 1 1.111-.757l4.12.78a.953.953 0 0 1 .757 1.11z"/><g fill="none" stroke="#000" stroke-miterlimit="10"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m29.915 14.262l4.864.954M21.903 12.69l4.93.967"/><circle cx="18.825" cy="55.057" r="5" stroke-width="2"/><circle cx="55.342" cy="55.057" r="5" stroke-width="2"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m27.529 17.318l-7.113 30.84l-1.591 6.899"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.908" d="M27.738 52.057h22.604v.583c0 1.657-1.223 3-2.73 3H30.467c-1.508 0-2.73-1.343-2.73-3v-.583z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29.85 14.25a1 1 0 0 1-.016.327l-.445 2.035a1 1 0 0 1-1.19.763l-1.06-.231a1 1 0 0 1-.762-1.19l.445-2.036c.02-.087.05-.17.088-.246m0 0a1 1 0 0 1 1.103-.517l1.058.232a1 1 0 0 1 .78.862m-7.197 25.668l5.084 12.14"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.908" d="m41.048 15.864l-.25 1.324a.954.954 0 0 1-1.112.758l-4.132-.783a.954.954 0 0 1-.758-1.112l.25-1.323a.954.954 0 0 1 1.112-.758l4.132.782a.954.954 0 0 1 .758 1.112zm-19.5-3.951l-.251 1.325a.953.953 0 0 1-1.11.757l-4.12-.78a.953.953 0 0 1-.758-1.11l.251-1.326a.953.953 0 0 1 1.111-.757l4.12.78a.953.953 0 0 1 .757 1.11z"/></g>`),
		g.Group(children),
	)
}