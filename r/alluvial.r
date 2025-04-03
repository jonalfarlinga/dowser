# Set the library path to include user libraries
.libPaths("C:/Users/nxl16/Documents/r/win-library/4.4") # Adjust the path accordingly

# Load necessary libraries
library(ggplot2)
library(ggalluvial)
library(colorspace)
library(ggrepel)

# Get command-line arguments
args <- commandArgs(trailingOnly = TRUE)
input_file <- "10_24_chart.csv"   # Path to the input CSV file
output_file <- "output.png"   # Path to save the output plot image

# Read the input CSV data
data <- read.csv(input_file, stringsAsFactors = FALSE)

# Check if the data is alluvial
if(!is_alluvia_form(data, key = c("Source", "Use", "Building"), value = "gals")) {
    stop("Data is not in a proper alluvial format.")
}

# Create the alluvial plot
p <- ggplot(data,
            aes(
                axis1 = Source,
                axis2 = Use,
                axis3 = Building,
                y = gals)) +
    geom_alluvium(aes(fill = Source), width = 1/12) +
    geom_stratum(width = 1/12, fill = "grey", color = "black") +
    #  geom_text(stat = "stratum", aes(label = after_stat(stratum)), size = 3) +
    scale_x_discrete(limits = c("Source", "Use", "Building"), expand = c(.05, .05)) +
    theme_minimal() +
    ggtitle("Water Flow from Sources to Uses to Buildings") +
    theme(plot.title = element_text(hjust = 0.5))

# Add non-overlapping labels using geom_text_repel with stat = "stratum"
p <- p + geom_text_repel(
  stat = "stratum",
  aes(label = after_stat(stratum)),
  size = 3,
  nudge_x = 0.05,
  direction = "y",
  segment.color = "grey50",
  max.overlaps = Inf
)

# Save the plot as a PNG file
ggsave(output_file, plot = p, width = 20, height = 12, dpi = 300)
