package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "os"
)

// @Summary Get policy
// @Description Returns system policy
// @Tags Policy&Privacy
// @Accept json
// @Produce json
// @Router /me [get]
func GetPolicy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Terms of Service\nLast Updated: 01.01.2025\n\nBy using the App, you agree to the following Terms of Service. Please read them carefully.\n\n1. Open-Source License\nThe App is released under the MIT License. You are free to use, modify, and distribute the App in accordance with the terms of the license. The source code is available at GitHub.\n\n2. No Warranty\nThe App is provided \"as is,\" without any warranties or guarantees of any kind, express or implied. The developers of the App are not liable for any damages or losses resulting from your use of the App.\n\n3. Your Responsibilities\nYou are responsible for:\n\nEnsuring the security of your device and data.\n\nBacking up your data to prevent loss.\n\nComplying with applicable laws and regulations when using the App.\n\n4. Third-Party Services\nIf you integrate the App with third-party services, you agree to comply with their terms of service and privacy policies. We are not responsible for any issues arising from the use of third-party services.\n\n5. Modifications to the App\nAs an open-source project, the App may be modified by you or other contributors. We are not responsible for any changes made by third parties.\n\n6. Termination\nYou may stop using the App at any time. We reserve the right to discontinue or modify the App at any time without notice.\n\n8. Contact Us\nIf you have any questions about these Terms of Service, please contact visit the project repository at GitHub.",
	})
}

// @Summary Get privacy
// @Description Returns system privacy
// @Tags Policy&Privacy
// @Accept json
// @Produce json
// @Router /me [get]
func GetPrivacy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Privacy Policy\nLast Updated: 01.01.2025\n\nThank you for using the App. This Privacy Policy explains how we handle your information when you use our open-source to-do app. Since the App is open-source, you are in control of your data and how it is used.\n\n1. Information We Do Not Collect\nThe App is designed to respect your privacy. We do not collect, store, or transmit any personal data or usage information. All data created or managed by you (e.g., tasks, reminders, or notes) is stored locally on your device unless you choose to sync it with a third-party service.\n\n2. Open-Source Nature\nThe App is open-source, meaning the source code is publicly available for review, modification, and distribution. You can inspect the code to verify that no data is being collected or transmitted without your consent.\n\n3. Third-Party Services\nIf you choose to integrate the App with third-party services (e.g., cloud storage or backup services), your data will be subject to the privacy policies of those services. We are not responsible for the practices of third-party services.\n\n4. Data Security\nSince the App does not collect or store your data on external servers, your data remains on your device. You are responsible for securing your device and any backups you create.\n\n5. Changes to This Policy\nWe may update this Privacy Policy from time to time. Any changes will be posted on this page, and the \"Last Updated\" date will be revised.\n\n6. Contact Us\nIf you have any questions about this Privacy Policy, please visit the project repository at GitHub.",
	})
}
