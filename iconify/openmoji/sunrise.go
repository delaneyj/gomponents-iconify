package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M4 47.292h64v20.787H4z"/><path fill="#e27022" d="M4 4.079h64v43.213H4z"/><path fill="#f4aa41" d="M68 4.079h-8l-16 31l4 3l17-20l3-3.231M4 17.641l2.019 1.913L24.55 38.145l3.75-3.308L9 4.079H4v13.562zM30.879 4.079l2.469 28.375l4.983.412l3.749-25.98l.486-2.807M68 34.719l-21.932 6.999l.707 4.95L68 44.947M4 34.899l21.933 6.999l-.707 4.95L4 45.127"/><path fill="#fcea2b" d="M21.48 47.292a15.68 15.68 0 1 1 29.04 0"/><path fill="#fff" d="M67.964 24.049L68 47.292H49.932l5.087-10.284l3.622-3.575l1.002-5.65l1.284-2.879l7.037-.855zm-63.927 0L4 47.292h17.48l-3.153-11.213l-3.479-3.988l.119-3.245l-3.894-3.942l-7.036-.855z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M8.48 32.188a7.753 7.753 0 0 1 9.848 4.82M5 24.026a7.753 7.753 0 0 1 9.848 4.82m39.627 11.737a7.753 7.753 0 0 1 8.332-7.127m-4.166-2.7a7.753 7.753 0 0 1 8.332-7.127M5.124 47.293h61.753m-46.387-3.33a15.68 15.68 0 1 1 30.775.253"/>`),
		g.Group(children),
	)
}