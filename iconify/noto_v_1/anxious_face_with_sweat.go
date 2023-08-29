package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnxiousFaceWithSweat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="notoV1AnxiousFaceWithSweat0" x1="63.79" x2="63.79" y1="65.667" y2="3.276" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#fcc21b"/><stop offset=".151" stop-color="#e5bf34"/><stop offset=".473" stop-color="#abb973"/><stop offset=".935" stop-color="#4fafd8"/><stop offset="1" stop-color="#42ade7"/></linearGradient><path fill="url(#notoV1AnxiousFaceWithSweat0)" d="M63.79 8.64C1.48 8.64 0 78.5 0 92.33c0 13.83 28.56 25.03 63.79 25.03c35.24 0 63.79-11.21 63.79-25.03c0-13.83-1.47-83.69-63.79-83.69z"/><path fill="#ed6c30" d="M86.34 90.53c0 13.17-10 18.72-22.34 18.72s-22.34-5.55-22.34-18.72s10-28.96 22.34-28.96s22.34 15.8 22.34 28.96z"/><path fill="#fff" d="M119.4 68.7c-1.51 4.55-4.85 7.13-7.49 7.67c-2.96.59-5.52-.1-7.2-1.95c-2.14-2.33-2.57-6.12-1.24-10.67c2.41-8.17 2.82-22.96 2.83-23.1c.01-.57.35-1.06.89-1.24c.52-.21 1.11-.05 1.49.37c12.9 14.41 12.27 24.3 10.72 28.92z"/><defs><path id="notoV1AnxiousFaceWithSweat1" d="M57.71 40.71c-.29-3.83-1.78-7.26-4.03-9.91c-11.8.13-22.02 4.42-26.37 6.55c-.46 1.8-.66 3.72-.51 5.69c.69 9.15 8.17 16.04 16.7 15.4s14.89-8.58 14.21-17.73z"/></defs><use fill="#fff" href="#notoV1AnxiousFaceWithSweat1"/><clipPath id="notoV1AnxiousFaceWithSweat2"><use href="#notoV1AnxiousFaceWithSweat1"/></clipPath><path fill="#2f2f2f" d="M54.68 33.48c.3 3.33-2.36 6.46-5.96 7.01c-3.6.55-6.76-1.71-7.07-5.03c-.32-3.33 2.35-6.46 5.95-7.01c3.59-.55 6.77 1.7 7.08 5.03" clip-path="url(#notoV1AnxiousFaceWithSweat2)"/><defs><path id="notoV1AnxiousFaceWithSweat3" d="M70.29 40.71c.29-3.83 1.78-7.26 4.03-9.91c11.8.13 22.02 4.42 26.37 6.55c.46 1.8.66 3.72.51 5.69c-.69 9.15-8.17 16.04-16.7 15.4s-14.89-8.58-14.21-17.73z"/></defs><use fill="#fff" href="#notoV1AnxiousFaceWithSweat3"/><clipPath id="notoV1AnxiousFaceWithSweat4"><use href="#notoV1AnxiousFaceWithSweat3"/></clipPath><path fill="#2f2f2f" d="M73.32 33.48c-.3 3.33 2.36 6.46 5.96 7.01c3.6.55 6.76-1.71 7.07-5.03c.32-3.33-2.35-6.46-5.95-7.01c-3.59-.55-6.77 1.7-7.08 5.03" clip-path="url(#notoV1AnxiousFaceWithSweat4)"/>`),
		g.Group(children),
	)
}