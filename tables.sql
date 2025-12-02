
CREATE TABLE `campaigns` (
  `id` bigint NOT NULL,
  `name` varchar(255) NOT NULL,
  `channel` enum('sms','whatsapp') NOT NULL,
  `status` enum('draft','scheduled','sending','sent','failed') NOT NULL DEFAULT 'draft',
  `base_template` text NOT NULL,
  `scheduled_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE `customers` (
  `id` bigint UNSIGNED NOT NULL,
  `phone` varchar(25) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `location` varchar(200) DEFAULT NULL,
  `preferred_product` varchar(100) DEFAULT NULL
);



CREATE TABLE `outbound_messages` (
  `id` int NOT NULL,
  `campaign_id` int NOT NULL,
  `customer_id` int NOT NULL,
  `status` enum('pending','sent','failed') NOT NULL DEFAULT 'pending',
  `rendered_content` text,
  `last_error` text,
  `retry_count` int NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;


--
-- FOREIGN KEYS -> Not set yet
-- 

INSERT INTO `campaigns` (`id`, `name`, `channel`, `status`, `base_template`, `scheduled_at`, `created_at`) VALUES
(1, 'Summer Sale 2025', 'sms', 'draft', 'Hi {first_name}, check out {preferred_product} in {location}!', '2025-06-01 10:00:00', '2025-12-02 16:20:42'),
(2, 'Winter Sale 2025', 'sms', 'draft', 'Hi {first_name}, check out {preferred_product} in {location}!', '2025-12-02 16:30:29', '2025-12-02 16:30:29');


INSERT INTO `customers` (`id`, `phone`, `first_name`, `last_name`, `location`, `preferred_product`) VALUES
(1, '0715401186', 'Hillary', 'Ngeno', 'Nairobi', 'Running shoes'),
(2, '', 'James', 'Bond', 'London', 'Whatsapp Services'),
(3, '', 'Ronald Reagan', 'J', 'London', 'Whatsapp Services');

