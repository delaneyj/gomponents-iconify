package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeamingFaceWithSmilingEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fcc21b" d="M63.89 9.64C1.58 9.64.1 79.5.1 93.33c0 13.83 28.56 25.03 63.79 25.03c35.24 0 63.79-11.21 63.79-25.03c0-13.83-1.47-83.69-63.79-83.69z"/><defs><path id="notoV1BeamingFaceWithSmilingEyes0" d="M63.89 98.06c23.15.05 40.56-12.97 41.19-29.05a232.203 232.203 0 0 1-82.38 0c.63 16.08 18.04 29.1 41.19 29.05z"/></defs><use fill="#fff" href="#notoV1BeamingFaceWithSmilingEyes0"/><clipPath id="notoV1BeamingFaceWithSmilingEyes1"><use href="#notoV1BeamingFaceWithSmilingEyes0"/></clipPath><g clip-path="url(#notoV1BeamingFaceWithSmilingEyes1)"><path fill="#2f2f2f" d="M78.05 108c-1.1 0-2-.9-2-2V61.07c0-1.1.9-2 2-2s2 .9 2 2V106a2 2 0 0 1-2 2z"/></g><g clip-path="url(#notoV1BeamingFaceWithSmilingEyes1)"><path fill="#2f2f2f" d="M92.21 108c-1.1 0-2-.9-2-2V61.07c0-1.1.9-2 2-2s2 .9 2 2V106a2 2 0 0 1-2 2z"/></g><g fill="#2f2f2f" clip-path="url(#notoV1BeamingFaceWithSmilingEyes1)"><path d="M63.89 108c-1.1 0-2-.9-2-2V61.07c0-1.1.9-2 2-2s2 .9 2 2V106c0 1.1-.9 2-2 2zm-14.17 0c-1.1 0-2-.9-2-2V61.07c0-1.1.9-2 2-2s2 .9 2 2V106a2 2 0 0 1-2 2zm-14.16 0c-1.1 0-2-.9-2-2V61.07c0-1.1.9-2 2-2s2 .9 2 2V106a2 2 0 0 1-2 2z"/></g><path fill="#2f2f2f" d="M64.01 100.56h-.25c-24.13 0-42.86-13.52-43.56-31.46c-.03-.76.29-1.49.86-1.98c.57-.5 1.33-.71 2.08-.57c26.82 4.84 54.67 4.84 81.5 0c.75-.14 1.51.08 2.08.57c.57.5.89 1.23.86 1.98c-.71 17.94-19.44 31.46-43.57 31.46zm-.13-5h.13c19.55 0 35.56-10.1 38.2-23.52a235.827 235.827 0 0 1-76.65 0c2.64 13.42 18.65 23.52 38.2 23.52h.12zM31.96 54.45a2.728 2.728 0 0 1-3.73.93a2.72 2.72 0 0 1-.96-3.71c.18-.31 4.6-7.62 14.37-7.62c9.78 0 14.2 7.31 14.39 7.62a2.735 2.735 0 0 1-2.36 4.11c-.92 0-1.83-.47-2.34-1.32c-.13-.22-3.12-4.96-9.69-4.96c-6.57-.01-9.54 4.74-9.68 4.95zm68.04.94c-.43.26-.91.38-1.37.38c-.94 0-1.85-.49-2.36-1.34c-.11-.2-3.08-4.94-9.66-4.94c-6.69 0-9.66 4.89-9.69 4.94a2.724 2.724 0 0 1-4.69-2.77c.18-.31 4.6-7.62 14.38-7.62c9.77 0 14.18 7.31 14.36 7.62c.76 1.3.32 2.97-.97 3.73z"/>`),
		g.Group(children),
	)
}