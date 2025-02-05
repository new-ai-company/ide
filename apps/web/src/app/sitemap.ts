import { MetadataRoute } from 'next'
import path from 'path'
import { getPageForSitemap } from '@/utils/sitemap'

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  let mergedSitemap: MetadataRoute.Sitemap = []

  const docsDirectory = path.join(
    process.env.NODE_ENV === 'production'
      ? path.join('.', 'src', 'app', '(docs)', 'docs')
      : path.join(process.cwd(), 'src', 'app', '(docs)', 'docs')
  )

  const docsPages = getPageForSitemap(
    docsDirectory,
    'https://e2b.dev/docs/',
    0.8
  )

  mergedSitemap = mergedSitemap.concat(docsPages)

  return mergedSitemap.sort((a, b) => a.url.localeCompare(b.url))
}
