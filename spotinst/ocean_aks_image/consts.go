package ocean_aks_image

import "https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"

const (
	Image commons.FieldName = "image"

	// Marketplace image.
	Marketplace commons.FieldName = "marketplace"
	Publisher   commons.FieldName = "publisher"
	Offer       commons.FieldName = "offer"
	SKU         commons.FieldName = "sku"
	Version     commons.FieldName = "version"
)
