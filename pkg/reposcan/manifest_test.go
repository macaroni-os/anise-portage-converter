/*
Copyright © 2021-2026 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package reposcan_test

import (
	"os"
	"path/filepath"

	. "github.com/macaroni-os/anise-portage-converter/pkg/reposcan"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reposcan Manifest", func() {

	Context("Parse Manifest", func() {
		It("Manifest 1", func() {

			m1 := `DIST lxd-5.6.tar.gz 16136915 BLAKE2B 9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209 SHA512 aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19
`

			manifest := ParseManifestContent([]byte(m1))

			Expect(manifest).ShouldNot(Equal(nil))
			Expect(len(manifest.Files)).Should(Equal(1))
			Expect(manifest.Files[0].Name).Should(Equal("lxd-5.6.tar.gz"))
			Expect(len(manifest.Files[0].Hashes)).Should(Equal(2))
			Expect(manifest.Files[0].Hashes["sha512"]).Should(Equal(
				"aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19"))

			Expect(manifest.Files[0].Hashes["blake2b"]).Should(Equal(
				"9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209"))

		})

		It("Manifest 2", func() {

			m1 := `DIST lxd-5.6.tar.gz 16136915 BLAKE2B 9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209 SHA512 aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19 MD5 fafsfsfsfsfsfsfsafsfsfsfsfsfsfsfsf
`

			manifest := ParseManifestContent([]byte(m1))

			Expect(manifest).ShouldNot(Equal(nil))
			Expect(len(manifest.Files)).Should(Equal(1))
			Expect(manifest.Files[0].Name).Should(Equal("lxd-5.6.tar.gz"))
			Expect(len(manifest.Files[0].Hashes)).Should(Equal(3))
			Expect(manifest.Files[0].Hashes["sha512"]).Should(Equal(
				"aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19"))

			Expect(manifest.Files[0].Hashes["blake2b"]).Should(Equal(
				"9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209"))

			Expect(manifest.Files[0].Hashes["md5"]).Should(Equal(
				"fafsfsfsfsfsfsfsafsfsfsfsfsfsfsfsf"))

		})

		It("Manifest 3", func() {

			tmpDir := os.TempDir()
			defer os.RemoveAll(tmpDir)

			m1 := `DIST lxd-5.6.tar.gz 16136915 BLAKE2B 9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209 SHA512 aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19 MD5 fafsfsfsfsfsfsfsafsfsfsfsfsfsfsfsf
`

			manifest := ParseManifestContent([]byte(m1))

			manifestFile := filepath.Join(tmpDir, "Manifest")

			errWrite := manifest.Write(manifestFile)

			manifest2, err := ParseManifest(manifestFile)

			Expect(errWrite).Should(BeNil())
			Expect(manifest).ShouldNot(Equal(nil))
			Expect(len(manifest.Files)).Should(Equal(1))
			Expect(manifest.Files[0].Name).Should(Equal("lxd-5.6.tar.gz"))
			Expect(len(manifest.Files[0].Hashes)).Should(Equal(3))
			Expect(manifest.Files[0].Hashes["sha512"]).Should(Equal(
				"aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19"))

			Expect(manifest.Files[0].Hashes["blake2b"]).Should(Equal(
				"9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209"))

			Expect(manifest.Files[0].Hashes["md5"]).Should(Equal(
				"fafsfsfsfsfsfsfsafsfsfsfsfsfsfsfsf"))

			Expect(manifest2.Files[0].Hashes["sha512"]).Should(Equal(
				"aae4e7675225fbe4592dfc63af66e4fe2c9198a0be65007b15bf5bda700725afa60929e1233392147a4d85a0058dd8a462a95910b945a776efef433997f4bc19"))

			Expect(manifest2.Files[0].Hashes["blake2b"]).Should(Equal(
				"9e8010c1b7d19e1220277f2bfa26d5bdfcc9f4cc9eab8408e1c3e2d42f41c6d87b616877890653ae0219470f2d037036b4a4f330d0d98edbd2e1c7d12b64f209"))

			Expect(manifest2.Files[0].Hashes["md5"]).Should(Equal(
				"fafsfsfsfsfsfsfsafsfsfsfsfsfsfsfsf"))

			Expect(err).Should(BeNil())
			Expect(manifest2).ShouldNot(Equal(nil))
			Expect(manifest).Should(Equal(manifest2))
		})

	})

})
