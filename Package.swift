// swift-tools-version: 5.9
import PackageDescription

#if os(Linux)
let revaultC: Target = .systemLibrary(name: "RevaultC", path: "CModule")
#else
let revaultC: Target = .binaryTarget(name: "RevaultC", url: "https://github.com/onepub-dev/reVault/releases/download/revault-api-v0.1.0/RevaultC.xcframework.zip", checksum: "d1862be6d86c9b16e355a20b1974685c2cf45d4686e59b8ddb6c974317e5ac83")
#endif

let package = Package(
    name: "RevaultAPI",
    products: [.library(name: "RevaultAPI", targets: ["RevaultAPI"])],
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", from: "1.29.0"),
    ],
    targets: [
        revaultC,
        .target(
            name: "RevaultAPI",
            dependencies: ["RevaultC", .product(name: "SwiftProtobuf", package: "swift-protobuf")],
            path: "Sources/RevaultAPI"
        ),
    ]
)
