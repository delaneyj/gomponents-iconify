package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenHandsDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiOpenHandsDarkSkinTone0" d="M6.523 24.298a2.51 2.51 0 1 0-3.565 3.47l9.417 9.675l.894.885l-5.153-5.343a2.51 2.51 0 1 0-3.564 3.47l6.443 6.62l5.094 5.04c4.389 3.963 10.743 2.954 14.705-1.434a13.94 13.94 0 0 0 3.553-9.238c-.087-3.02-.729-10.864-.729-10.864c-.35-1.82-2.595-2.925-2.75-1.946l-2.625 7.3l-2.733-2.694l2.733 2.694l-2.733-2.694l-10.96-11a2.51 2.51 0 1 0-3.564 3.47l3.47 3.565l4.77 4.941L8.74 19.376a2.51 2.51 0 1 0-3.565 3.47l9.552 9.718m51.097-8.266a2.51 2.51 0 1 1 3.564 3.47l-9.417 9.675l-.893.885l5.152-5.343a2.51 2.51 0 1 1 3.565 3.47l-6.443 6.62l-5.094 5.04c-4.39 3.963-10.744 2.954-14.705-1.434A13.94 13.94 0 0 1 38 37.443c.087-3.02.729-10.864.729-10.864c.349-1.82 2.595-2.925 2.75-1.946l2.625 7.3l2.733-2.694l-2.733 2.694l2.733-2.694l10.96-11a2.51 2.51 0 1 1 3.564 3.47l-3.47 3.565l-4.771 4.942l10.487-10.84a2.51 2.51 0 1 1 3.565 3.47l-9.553 9.718"/></defs><use href="#openmojiOpenHandsDarkSkinTone0" fill="#6a462f" stroke="#6a462f"/><use href="#openmojiOpenHandsDarkSkinTone0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`),
		g.Group(children),
	)
}